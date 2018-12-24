<template>
  <div id="app" style="position:absolute; left:0; top:0; width:100%;height:100%;">
    <el-container>
      <el-header class="app-el-header">
        <el-menu class="el-menu" mode="horizontal" :default-active="activeIndex" @select="changeRoute">
          <el-menu-item index="/me">me</el-menu-item>
          <el-menu-item index="/articles">articles</el-menu-item>
          <el-menu-item v-show="notSignedIn" class="menu-right" index="/signin">sign in</el-menu-item>
          <el-menu-item v-show="notSignedIn && openSignup" class="menu-right" index="/signup">sign up</el-menu-item>
          <el-submenu class="menu-right" v-show="!notSignedIn" :show-timeout="1" :hide-timeout="1" index="#">
            <template slot="title">{{ nickname }}</template>
            <el-menu-item class="subMenuItem" index="/profile">profile</el-menu-item>
            <el-menu-item class="subMenuItem" index="/careers">careers</el-menu-item>
            <el-menu-item class="subMenuItem" index="/educations">educations</el-menu-item>
            <el-menu-item class="subMenuItem" index="/signout" @click="signout">sign out</el-menu-item>
          </el-submenu>
          <el-menu-item v-show="!notSignedIn" class="menu-right" index="/new">new</el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
      <router-view/>
      </el-main>
      <div class="line"></div>
      <el-footer>
        Sang Hongwei @ 2018
      </el-footer>
    </el-container>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { mapMutations } from 'vuex'

export default {
  name: "app",
  data() {
    return {
      route: {
        "/me": {name: 'me'},
        "/signup": {name: 'signup'},
        "/signin": {name: 'signin'},
        "/signout": {name: 'signout'},
        "/new": {name: 'new'},
        "/profile": {name: 'profile'},
        "/careers": {name: 'careers'},
        "/educations": {name: 'educations'},
        "/articles": {name: 'articles'},
      }
    }
  },

  computed: {
    expired() {
      return this.$store.state.expired
    },
    ...mapGetters([
      'activeIndex',
      'notSignedIn',
      'openSignup',
      'nickname',
    ])
  },

  methods: {
    changeRoute(key, keyPath) {
      this.$router.push(this.route[key])
    },

    signout() {
      this.SIGN_OUT()
      this.$notify({
        title: 'Sign out success.',
        position: 'bottom-right',
        type: 'success'
      })
      this.$router.push("/")
    },
  
    ...mapMutations([
      'ACTIVE_INDEX',
      'SIGN_OUT',
    ])
  },

  watch:{
    $route(to,from){
      this.ACTIVE_INDEX(to.path)
    },
    expired(now, old) {
      if(now == true){
        this.SIGN_OUT()
        this.$message({
          type: 'warning',
          message: 'Your token has expired, login first.'
        });
        setTimeout(function(){
          router.push("/signin")
        }, 1000)
      }
    }
  },

  mounted() {
    this.ACTIVE_INDEX(this.$route.path)
  },
}
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 100%;
}

.line {
  height: 1px;
  background-color: #e0e6ed;
}

.el-main {
  color: #333;
  text-align: center;
  padding: 0 !important;
}

.el-container {
  min-height: 100%;
}

.el-header,
.el-footer {
  color: #333;
  text-align: center;
  line-height: 60px;
  padding: 0 !important;
}

.menu-right {
  float: right !important;
}

.el-menu-item .subMenuItem{
  font-size: 18px !important;
}

.subMenuItem{
  text-align: center;
}
</style>
