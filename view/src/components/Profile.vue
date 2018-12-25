<template>
  <div class="profile">
    <el-row>
      <el-col :span="18" :offset="3">
        <el-form :model="form" :rules="profileRules" ref="form" label-width="160px">
          <el-row>
            <el-col :span="16">
              <el-row>
                <el-col :span="24">
                  <el-row style="text-align: right">
                    <span style="margin-right:10px">Upload your picture here</span>
                    <el-button @click="toggleShow" style="margin:22px 0;">Upload Image</el-button></el-row>
                  <el-form-item label="Nickname" prop="nickname">
                    <el-input v-model="form.nickname" placeholder="Change your nickname"></el-input>
                  </el-form-item>
                  <el-form-item label="Realname" prop="name">
                    <el-input v-model="form.name" placeholder="Enter your real name"></el-input>
                  </el-form-item>
                  <el-form-item label="Gender" prop="gender" style="text-align: left">
                    <el-radio v-model="form.gender" label="1">
                      Male<i class="el-icon-ali-male_icon" style="margin-left:3px;"></i>
                    </el-radio>
                    <el-radio v-model="form.gender" label="2">
                      Female<i class="el-icon-ali-female_icon" style="margin-left:3px;"></i>
                    </el-radio>
                    <el-radio v-model="form.gender" label="0">
                      Secret<i class="el-icon-ali-yinsi" style="margin-left:3px;"></i>
                    </el-radio>
                  </el-form-item>
                  <el-form-item label="Contact" prop="email">
                    <el-input type="email" v-model="form.email" placeholder="Enter your Email"></el-input>
                  </el-form-item>
                  <el-form-item label="Birthday" prob="birthday" style="text-align: left">
                    <el-date-picker v-model="form.birthday" type="date" placeholder="Choose your birthday"></el-date-picker>
                  </el-form-item>
                  <el-form-item label="Hometown" prop="hometown">
                    <el-input v-model="form.hometown" placeholder="Enter your hometown"></el-input>
                  </el-form-item>
                  <el-form-item label="City" prop="city">
                    <el-input v-model="form.city" placeholder="Enter the city where you live now"></el-input>
                  </el-form-item>
                  <el-form-item label="Me">
                    <el-input type="textarea" v-model="form.me" rows="2" placeholder="Yourself"></el-input>
                  </el-form-item>
                </el-col>
              </el-row>
            </el-col>
            <el-col :span="8">
              <img :src="imgDataUrl" style="margin: 40px 10px;">
            </el-col>
          </el-row>
          <el-button type="primary" @click="submitForm('form')" style="margin-bottom:22px;">Submit</el-button></el-form>
      </el-col>
      <my-upload
        field="img"
        @crop-success="cropSuccess"
        @crop-upload-success="cropUploadSuccess"
        @crop-upload-fail="cropUploadFail"
        v-model="show"
        :width="300"
        :height="400"
        url="http://www.sanghongwei.com:8080/user/image"
        :params="form"
        :headers="headers"
        :noCircle=true
        :withCredentials=true
        img-format="png">
      </my-upload>
    </el-row>
  </div>
</template>

<script>
import {mapActions} from "vuex"
//import 'babel-polyfill'
import myUpload from 'vue-image-crop-upload'

export default {
  name: 'Profile',
  components: {
    'my-upload': myUpload
  },

  data: function() {
    var validateText = (rule, value, callback) => {
      if (!this.checkSpecificKey(value)) {
        callback(new Error('不能包含特殊字符'))
      } else {
        callback()
      }
    }

    return {
      show: false,
      headers: {
        "X-USER-TOKEN": '',
      },
      imgDataUrl: '',
      form: {
        id: '',
        nickname: '',
        me: '',
        name: '',
        email: '',
        gender: '0',
        hometown: '',
        city: '',
        birthday: null,
      },
      profileRules: {
        nickname: [
          { required: true, message: '请输入昵称', trigger: 'blur' },
          { validator: validateText, trigger: ['blur', 'change']},
          { max: 16, message: '昵称不能大于16个字符', trigger: 'blur'}
        ],
        name: [
          { validator: validateText, trigger: ['blur', 'change']},
          { max: 16, message: '姓名不能大于16个字符', trigger: 'blur'}
        ],
        email: [
          { type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] },
          { min: 5, max: 64, message: '长度必须在5到64之间', trigger: 'blur'}
        ],
        hometown: [
          { max: 32, message: 'Hometown不能大于16个字符', trigger: 'blur'}
        ],
        city: [
          { max: 32, message: 'City不能大于16个字符', trigger: 'blur'}
        ],
      }
    };
  },
  
  methods: {
    toggleShow() {
      this.show = !this.show;
    },

    cropSuccess(imgDataUrl, field){
      console.log('-------- crop success --------');
      this.imgDataUrl = imgDataUrl;
    },

    cropUploadSuccess(jsonData, field){
      console.log('-------- upload success --------');
      console.log(jsonData);
      console.log('field: ' + field);
    },

    cropUploadFail(status, field){
      console.log('-------- upload fail --------');
      console.log(status);
      console.log('field: ' + field);
    },

    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.EDIT_USER_PROFILE(this.form).then((res) => {
            this.$notify({
              title: 'Edit profile success',
              position: 'bottom-right',
              type: 'success'
            })
            router.push({
              path: '/me'
            })
          }, (error) => {
            this.$notify.error({
              title: 'Sign in failed.',
              message: error.response.data.msg,
              position: "top-right"
            })
          })
        }
      })
    },

    ...mapActions([
      'GET_USER_PROFILE',
      'EDIT_USER_PROFILE',
      'GET_USER_PROFILE_IMAGE',
    ])
  },

  mounted(){
    this.headers["X-USER-TOKEN"] = this.$store.state.token
    this.form.id = this.$store.state.user_id.toString()
    this.GET_USER_PROFILE(this.$store.state.user_id.toString()).then((res) => {
      this.form.nickname = res.data.nickname
      this.form.me = res.data.me
      this.form.name = res.data.name
      this.form.gender = res.data.gender.toString()
      this.form.email = res.data.email
      this.form.birthday = res.data.birthday
      this.form.hometown = res.data.hometown
      this.form.city = res.data.city
    }, (error) => {
      this.$notify.error({
        title: 'Load user profile failed.',
        message: error.response.data.msg,
        position: "top-right"
      })
    }),
    this.GET_USER_PROFILE_IMAGE(this.$store.state.ownerID).then((res) => {
      this.imgDataUrl = 'data:image/jpeg;base64,' + res.data.img
    }, (error) => {
      this.$notify.error({
        title: 'Load user profile img failed.',
        message: error.response.data.msg,
        position: "top-right"
      })
    })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
