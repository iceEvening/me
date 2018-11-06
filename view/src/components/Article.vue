<template>
  <div class="article">
    <el-row>
      <el-col :span="12" :offset="6" id="title">
        <h3>{{ article.title }}</h3>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12" :offset="6" id="title">
        {{ article.nickname }}<br>{{ article.create_time }}
      </el-col>
    </el-row>
    <el-row type="flex" justify="center">
      <el-col :span="16">
       <div id="article_content" v-html="article.html">
       </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import {mapActions} from "vuex"
import router from '../router'

export default {
  name: 'Article',
  
  data() {
    return {
      article: {
        title: 'Loading...',
        nickname: 'unknow',
        create_time : '1970年01月01日 00:00:00',
        html: 'Loading...',
      }
    }
  },

  methods: {
    ...mapActions([
      'ARTICLE',
    ])
  },

  mounted() {
    this.ARTICLE(this.$route.params.id).then((res) => {
      console.log('Load article success!')
      this.article.title = res.data.title
      this.article.nickname = res.data.nickname
      this.article.create_time = res.data.create_time
      this.article.html = res.data.html
    }, (error) => {
      this.$notify.error({
        title: 'Load article failed.',
        message: error.response.data.msg,
        position: "top-right"
      })
    })
  },
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

#article_content {
  text-align: left;
}

</style>
