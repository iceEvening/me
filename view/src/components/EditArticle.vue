<template>
  <div class="edit_article">
    <el-row>
      <el-col :span="12" :offset="6" id="title" style="margin-bottom:10px;margin-top:10px;">
        <span>New article</span>
      </el-col>
    </el-row>
    <el-row type="flex" justify="center">
      <el-col :span="20">
        <el-form :model="article" :rules="newArticleRules" ref="form" label-width="60px">
          <el-form-item label="Title" prop="title">
            <el-input v-model="article.title" placeholder="Please input the title..."></el-input>
          </el-form-item>
          <mavon-editor
            @change="getHtml"
            v-model="article.markdown"
            :toolbars="toolbars"
            :codeStyle="codeStyle"
            :boxShadow=false
            style="height: 500px"/>
          <el-form-item id="submit">            
              <el-button type="primary" @click="submitForm('form')" style="margin-top:20px">提交</el-button>
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
  name: 'EditArticle',
  data() {
    return {
      article: {
        title: '',
        markdown: '',
        html: '',
      },
      newArticleRules: {
        title: [
          { required: true, message: '请输入Title', trigger: 'blur' },
          { min: 1, max: 64, message: '长度必须在1到64之间', trigger: 'blur'}
        ],
      },
      codeStyle: 'monokai',
      toolbars: {
        bold: true, // 粗体
        italic: true, // 斜体
        header: true, // 标题
        underline: true, // 下划线
        strikethrough: true, // 中划线
        mark: false, // 标记
        superscript: false, // 上角标
        subscript: false, // 下角标
        quote: true, // 引用
        ol: true, // 有序列表
        ul: true, // 无序列表
        link: true, // 链接
        imagelink: true, // 图片链接
        code: true, // code
        table: true, // 表格
        fullscreen: true, // 全屏编辑
        readmodel: true, // 沉浸式阅读
        htmlcode: true, // 展示html源码
        help: true, // 帮助
        /* 1.3.5 */
        undo: true, // 上一步
        redo: true, // 下一步
        trash: true, // 清空
        save: false, // 保存（触发events中的save事件）
        /* 1.4.2 */
        navigation: true, // 导航目录
        /* 2.1.8 */
        alignleft: true, // 左对齐
        aligncenter: true, // 居中
        alignright: true, // 右对齐
        /* 2.2.1 */
        subfield: true, // 单双栏模式
        preview: true, // 预览
      }
    }
  },
  methods: {
    getHtml(value, render) {
      this.article.html = render
    },
    
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.EDIT_ARTICLE({
            'id': this.$route.params.id,
            'title': this.article.title,
            'markdown': this.article.markdown,
            'html': this.article.html,
            'user_id': this.$store.state.user_id.toString(),
            'nickname': this.$store.state.nickname,
          }).then((res) => {
            console.log('Edit article success!')
            this.$notify({
              title: 'Success',
              message: 'Create article success!',
              position: 'bottom-right',
              type: 'success'
            })
            router.push({
              path: '/article/' + this.$route.params.id
            })
          }, (error) => {
            this.$notify.error({
              title: 'Edit article failed.',
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
      'ARTICLE',
      'EDIT_ARTICLE',
    ])
  },

  mounted() {
    this.ARTICLE(this.$route.params.id).then((res) => {
      console.log('Load article success!')
      this.article.title = res.data.title
      this.article.markdown = res.data.markdown
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
#submit {
  margin-left: -60px;
}
#title {
  line-height: 0;
  padding-top: 20px;
  padding-bottom: 20px;
}

</style>
