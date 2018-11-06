<template>
  <div class="careers">
    <el-row>
      <el-col :span="12" :offset="6" id="title" style="margin-bottom:20px;margin-top:20px;">
        <span>Edit careers</span>
      </el-col>
    </el-row>
    <el-form :model="careers" ref="careers" label-width="90px">
      <el-row v-for="(career, i) in careers.list" :key="i" type="flex" justify="center">
        <el-col :span="10">
          <!-- card start-->
          <el-card class="career-card" shadow="hover">
            <!-- card header -->
            <div slot="header" class="clearfix">
              <i class="el-icon-date" style="padding-right: 3px"></i>
              <span>Career {{ i+1 }}</span>
              <el-checkbox v-model="career.isCurrent" style="padding-left: 10px">current</el-checkbox>
              <el-button style="float: right; padding: 3px 0" type="text" @click="add">Add</el-button>
              <i class="el-icon-circle-plus-outline" style="float: right; padding: 3px 0"></i>
              <el-button style="float: right; padding: 3px 0; padding-right: 3px;" type="text" @click="del(i, career.id)">Delete</el-button>
              <i class="el-icon-delete" style="float: right; padding: 3px 0"></i>
            </div>
            <!-- card content -->
            <div>
              <el-form-item label="Company" :prop="'list.' + i + '.company'" :rules="careersRules.company">
                <el-input v-model="career.company" placeholder="Please enter your company's name"></el-input>
              </el-form-item>
              <el-form-item label="Department" :prop="'list.' + i + '.department'" :rules="careersRules.department">
                <el-input v-model="career.department" placeholder="Please enter your department in company"></el-input>
              </el-form-item>
              <el-row type="flex" justify="left">
                <el-form-item label="Post" :prop="'list.' + i + '.post'" :rules="careersRules.post">
                  <el-input v-model="career.post" placeholder="Position in company"></el-input>
                </el-form-item>
                <el-form-item label="Rank" :prop="'list.' + i + '.rank'" :rules="careersRules.rank">
                  <el-input v-model="career.rank" placeholder="Rank in company"></el-input>
                </el-form-item>
              </el-row>
              <el-form-item label="Description" :prop="'list.' + i + '.description'" :rules="careersRules.description">
                <el-input type="textarea" autosize v-model="career.description" placeholder="Please enter your job description"></el-input>
              </el-form-item>
              <el-row type="flex" justify="left">
                <el-form-item label="Start time" :prop="'list.' + i + '.start'" :rules="careersRules.start">
                  <el-date-picker v-model="career.start" type="month" placeholder="Month" style="width: 178px;"></el-date-picker>
                </el-form-item>
                <el-form-item v-show="!career.isCurrent" label="End time" :prop="'list.' + i + '.end'" :rules="careersRules.end">
                  <el-date-picker v-model="career.end" type="month" placeholder="Month" style="width: 178px;"></el-date-picker>
                </el-form-item>
              </el-row>
            </div>
          </el-card>
          <!-- card end-->
        </el-col>
      </el-row>
      <el-row type="flex" justify="center">
        <el-col style="margin-bottom:20px; margin-left: -90px;">
          <el-form-item id="submit">            
            <el-button type="primary" @click="submitForm('careers')">提交</el-button>
            <el-button @click="resetForm('careers')">重置</el-button>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
  </div>
</template>

<script>
import {mapActions} from "vuex";

var career = {
  id : '0',
  company: '',
  department : '',
  post: '',
  rank: '',
  description: '',
  start: null,
  end: new Date(),
  isCurrent: false,
}

export default {
  name: 'Careers',

  data() {
    return {
      careers: {
        list: []
      },
      careersRules: {
        company: [
          { required: true, message: 'Please input company', trigger: 'blur'},
          { min: 2, max: 32, message: 'Text length must between 2 and 32', trigger: 'blur'}
        ],
        department: [
          { required: true, message: 'Please input department', trigger: 'blur'},
          { min: 2, max: 32, message: 'Text length must between 2 and 32', trigger: 'blur'}
        ],
        post: [
          { required: true, message: 'Please input post', trigger: 'blur'},
          { min: 2, max: 32, message: 'Text length must between 2 and 32', trigger: 'blur'}
        ],
        rank: [
          { required: true, message: 'Please input job level', trigger: 'blur'},
          { min: 2, max: 16, message: 'Text length must between 2 and 16', trigger: 'blur'}
        ],
        description: [
          { required: true, message: 'Please input department', trigger: 'blur'},
          { min: 2, max: 8192, message: 'Text length must between 2 and 8192', trigger: 'blur'}
        ],
        start: [
          { required: true, message: 'Please input department', trigger: 'blur'},
        ],
        end: [
          { required: true, message: 'Please input department', trigger: 'blur'},
        ],
      }
    }
  },

  methods: {
    add() {
      this.careers.list.push(Object.assign({},career))
      console.log(this.careers)
    },
    del(i, id) {
      this.$confirm('Delete this career?', 'Alert', {
        confirmButtonText: 'Delete',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(() => {
        this.DEL_CAREER({
          'id': id.toString(),
          'user_id': this.$store.state.user_id.toString(),
        }).then((res) => {
          this.careers.list.splice(i, 1)
          this.$message({
            type: 'success',
            message: 'delete success!'
          });
        }, (error) => {
          this.$notify.error({
            title: 'Delete career failed.',
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
    submitForm() {
      this.CAREERS({
        'id': this.$store.state.user_id.toString(),
        'list': this.careers.list}).then((res) => {
        this.$notify({
          title: 'Success',
          message: 'Edit careers success!',
          position: 'bottom-right',
          type: 'success'
        })
      }, (error) => {
        this.$notify.error({
          title: 'Edit careers failed.',
          message: error.response.data.msg,
          position: "top-right"
        })
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    ...mapActions([
      'CAREERS',
      'GET_CAREERS',
      'DEL_CAREER',
    ])
  },

  mounted() {
    this.GET_CAREERS(this.$store.state.user_id.toString()).then((res) => {
      console.log('Get careers success!')
      if (res.data.careers.length > 0) {
        var careerList = []
        res.data.careers.forEach(function(eachCareer){
          var c = Object.assign({}, career)
          c.id = eachCareer.ID.toString()
          c.company = eachCareer.Company
          c.department = eachCareer.Department
          c.post = eachCareer.Post
          c.rank = eachCareer.JobLevel
          c.description = eachCareer.JobDesc
          c.start = eachCareer.StartTime
          c.end = eachCareer.EndTime
          c.isCurrent = eachCareer.IsCurrent
          careerList.push(c)
        })
        this.careers.list = careerList
      } else {
        this.careers.list = [Object.assign({},career)]
      }
    }, (error) => {
      this.$notify.error({
        title: 'Get careers failed.',
        message: error.response.data.msg,
        position: "top-right"
      })
    })
  },
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

.career-card {
  text-align: left;
  margin-bottom: 30px;
}
</style>
