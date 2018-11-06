<template>
  <div class="articles">
    <el-row style="padding: 40px 0;">
      <el-col :span="18" :offset="3">
        <el-row>
          <el-col :span="0">
            
          </el-col>
          <el-col :span="24" style="text-align: left; padding:0 10px;">
            <el-row>
              <el-col :span="24" style="margin-bottom:20px;margin-top:20px;text-align:center;">
                <span>Articles</span>
              </el-col>
            </el-row>
            <el-row style="margin:30px 0;">
              <el-col :span="24">
                <el-table
                  :show-header="false"
                  :default-sort="sort"
                  :data="articles.list"
                  style="width: 100%">
                  <el-table-column
                    type="index"
                    width="50">
                  </el-table-column>
                  <el-table-column
                    prop="title"
                    label="Title">
                    <template slot-scope="scope">
                      <a :href="'article/'+scope.row.id">{{ scope.row.title }}</a>
                    </template>
                  </el-table-column>
                  <el-table-column
                    prop="nickname"
                    label="Nickname"
                    width="150">
                  </el-table-column>
                  <el-table-column
                    prop="update"
                    label="Update"
                    width="200">
                  </el-table-column>
                  <el-table-column
                    v-if="isEdit"
                    width="160"
                    label="operate">
                    <template slot-scope="scope">
                      <el-button size="mini" @click="edit(scope.row)">edit</el-button>
                      <el-button size="mini" type="danger" @click="del(scope.row)">delete</el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </el-col>
            </el-row>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import {mapActions} from "vuex"
import router from '../router'

var article = {
  id : '0',
  title: '',
  nickname : '',
  create: '',
  update: '',
}

export default {
  name: 'Article',
  
  data() {
    return {
      articles: {
        list: []
      },

      sort: {
        prop: 'update',
        order: 'descending'
      },

      isEdit: false,

      style: {
        cursor: 'pointer'
      }
    }
  },

  methods: {
    titleClick(row, column, cell, event) {
      if(column.label === 'Title') {
        this.$router.push("/article/" + row.id)
      }
    },

    changeMouse({row, column, rowIndex, columnIndex}) {
      if(column.label === 'Title') {
        return {cursor: 'pointer'}
      }
    },

    edit(row) {
      this.$router.push("/editarticle/" + row.id)
    },

    del(row) {
      this.$confirm('Delete this article?', 'Alert', {
        confirmButtonText: 'Delete',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(() => {
        this.DEL_ARTICLE({
          'id': row.id.toString(),
          'user_id': this.$store.state.user_id.toString(),
        }).then((res) => {
          this.$message({
            type: 'success',
            message: 'delete success!'
          });
          setTimeout(function(){
            router.go(0)
          }, 1000)
        }, (error) => {
          this.$notify.error({
            title: 'Delete article failed.',
            message: error.response.data.msg,
            position: "top-right"
          })
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: 'canceled'
        });          
      });
    },
  
    ...mapActions([
      'ARTICLES',
      'DEL_ARTICLE'
    ])
  },

  mounted() {
    if(this.$store.state.user_id == this.$store.state.ownerID) {
      this.isEdit = true
    }
    this.ARTICLES(this.$store.state.ownerID).then((res) => {
      if (res.data.articles.length > 0) {
        var articles = []
        res.data.articles.forEach(function(eachArticle){
          var a = Object.assign({}, article)
          a.id = eachArticle.ID
          a.title = eachArticle.Title
          a.nickname = eachArticle.NickName
          a.create = eachArticle.CreatedAt
          a.update = new Date(eachArticle.UpdatedAt).toLocaleString()
          articles.push(a)
        })
        this.articles.list = articles
      } else {
        
      }
    }, (error) => {
      this.$notify.error({
        title: 'Get articles failed.',
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

.title{
  cursor:pointer;
}

a:link {
  text-decoration: none;
  color:#606266;
}

a:visited {
  text-decoration: none;
  color: #606266;
}

</style>
