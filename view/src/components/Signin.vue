<template>
  <div class="me">
    <el-row>
      <el-col :span="12" :offset="6" id="title" style="margin-bottom:20px;margin-top:20px;">
        <span>Sign in</span>
      </el-col>
    </el-row>
    <el-row type="flex" justify="center">
      <el-col :xs="20" :sm="20" :md="16" :lg="8" :xl="6">
        <el-form :model="form" :rules="signinRules" ref="form" label-width="60px">
          <el-form-item label="邮箱" prop="email">
            <el-input type="email" v-model="form.email" @keyup.enter.native="submitForm('form')" placeholder="请输入Email"></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input type="password" v-model="form.password" autocomplete="off" @keyup.enter.native="submitForm('form')" placeholder="请输入密码" ></el-input>
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
  name: 'Signin',
  data() {
    return {
      form: {
        email: '',
        password: '',
      },
      signinRules: {
        email: [
          { required: true, message: '请输入邮箱地址', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] },
          { min: 5, max: 64, message: '长度必须在5到64之间', trigger: 'blur'}
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 8, max: 32, message: '长度必须在8到32之间', trigger: 'blur'}
        ],
      }
    }
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.SIGN_IN(this.form).then((res) => {
            console.log('Successfully sign in!')
            console.log(this.$store.state.token)
            this.$notify({
              title: 'Hello ' + this.$store.state.email + '.',
              message: 'Welcome back.',
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

    resetForm(formName) {
      this.$refs[formName].resetFields()
    },

    ...mapActions([
      'SIGN_IN',
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
