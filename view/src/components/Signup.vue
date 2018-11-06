<template>
  <div class="me">
    <el-row>
      <el-col :span="12" :offset="6" id="title" style="margin-bottom:20px;margin-top:20px;">
        <span>Sign up</span>
      </el-col>
    </el-row>
    <el-row type="flex" justify="center">
      <el-col :xs="20" :sm="20" :md="16" :lg="8" :xl="6">
        <el-form :model="form" :rules="rules" ref="form" label-width="60px">
          <el-form-item label="邮箱" prop="email">
            <el-input type="email" v-model="form.email" placeholder="Email将作为您的登录凭证，不可更改"></el-input>
          </el-form-item>
          <el-form-item label="昵称" prop="nickname">
            <el-input v-model="form.nickname" placeholder="您的昵称，可随时修改"></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input type="password" v-model="form.password" placeholder="请输入密码" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="确认" prop="chkPassword">
            <el-input type="password" v-model="form.chkPassword" placeholder="请再次输入密码" autocomplete="off" @keyup.enter.native="submitForm('form')"></el-input>
          </el-form-item>
          <el-form-item id="submit">            
              <el-button type="primary" @click="submitForm('form')">提交</el-button>
              <el-button @click="resetForm('form')">重置</el-button>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import {mapActions} from "vuex";
import router from '../router'

export default {
  name: 'Signup',
  data() {
    var validateText = (rule, value, callback) => {
      if (!this.checkSpecificKey(value)) {
        callback(new Error('不能包含特殊字符'))
      } else {
        callback()
      }
    }

    var validatePwd = (rule, value, callback) => {
      if (this.form.chkPassword !== '') {
        this.$refs.form.validateField('chkPassword')
      }
      callback()
    }

    var validateChkPwd = (rule, value, callback) => {
      if (value !== this.form.password) {
        callback(new Error('两次输入密码不一致'))
      } else {
        callback()
      }
    }
  
    return {
      form: {
        email: '',
        nickname: '',
        password: '',
      },
      rules: {
        email: [
          { required: true, message: '请输入邮箱地址', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] },
          { min: 5, max: 64, message: '长度必须在5到64之间', trigger: 'blur'}
        ],
        nickname: [
          { required: true, message: '请输入昵称', trigger: 'blur' },
          { validator: validateText, trigger: ['blur', 'change']},
          { max: 16, message: '昵称不能大于16个字符', trigger: 'blur'}
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { validator: validatePwd, trigger: 'blur' },
          { min: 8, max: 32, message: '长度必须在8到32之间', trigger: 'blur'}
        ],
        chkPassword: [
          { required: true, message: '请再次输入密码', trigger: 'blur' },
          { validator: validateChkPwd, trigger: 'blur'}
        ]
      }
    }
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.SIGN_UP(this.form).then((res) => {
            this.$notify({
              title: 'Sign up success.',
              message: 'Sign in your account now.',
              position: 'bottom-right',
              type: 'success'
            })
            router.push({
              path: '/signin'
            })
          }, (error) => {
            this.$notify.error({
              title: 'Sign up failed.',
              message: error.response.data.msg,
              position: "top-right"
            });
          })
        } else {
          return false
        }
      })
    },

    resetForm(formName) {
      this.$refs[formName].resetFields()
    },

    ...mapActions([
      'SIGN_UP',
    ])
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
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
#submit {
  margin-left: -60px;
}
#title {
  line-height: 0;
  padding-top: 20px;
  padding-bottom: 20px;
}

</style>
