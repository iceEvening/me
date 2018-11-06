<template>
  <div class="educations">
    <el-row>
      <el-col :span="12" :offset="6" id="title" style="margin-bottom:20px;margin-top:20px;">
        <span>Edit educations</span>
      </el-col>
    </el-row>
    <el-form :model="educations" ref="educations" label-width="90px">
      <el-row v-for="(education, i) in educations.list" :key="i" type="flex" justify="center">
        <el-col :span="10">
          <!-- card start-->
          <el-card class="education-card" shadow="hover">
            <!-- card header -->
            <div slot="header" class="clearfix">
              <i class="el-icon-date" style="padding-right: 3px"></i>
              <span>Education {{ i+1 }}</span>
              <el-checkbox v-model="education.isCurrent" style="padding-left: 10px">current</el-checkbox>
              <el-button style="float: right; padding: 3px 0" type="text" @click="add">Add</el-button>
              <i class="el-icon-circle-plus-outline" style="float: right; padding: 3px 0"></i>
              <el-button style="float: right; padding: 3px 0; padding-right: 3px;" type="text" @click="del(i, education.id)">Delete</el-button>
              <i class="el-icon-delete" style="float: right; padding: 3px 0"></i>
            </div>
            <!-- card content -->
            <div>
              <el-form-item label="School" :prop="'list.' + i + '.school'" :rules="educationsRules.school">
                <el-input v-model="education.school" placeholder="Please enter your school's name"></el-input>
              </el-form-item>
              <el-form-item label="Major" :prop="'list.' + i + '.major'" :rules="educationsRules.major">
                <el-input v-model="education.major" placeholder="Please enter your major at school"></el-input>
              </el-form-item>
              <el-form-item label="Degree" :prop="'list.' + i + '.degree'" :rules="educationsRules.degree">
                <el-input v-model="education.degree" placeholder="Degree at school"></el-input>
              </el-form-item>
              <el-row type="flex" justify="left">
                <el-form-item label="Start time" :prop="'list.' + i + '.start'" :rules="educationsRules.start">
                  <el-date-picker v-model="education.start" type="month" placeholder="Month" style="width: 178px;"></el-date-picker>
                </el-form-item>
                <el-form-item v-show="!education.isCurrent" label="End time" :prop="'list.' + i + '.end'" :rules="educationsRules.end">
                  <el-date-picker v-model="education.end" type="month" placeholder="Month" style="width: 178px;"></el-date-picker>
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
            <el-button type="primary" @click="submitForm('educations')">提交</el-button>
            <el-button @click="resetForm('educations')">重置</el-button>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
  </div>
</template>

<script>
import {mapActions} from "vuex";

var education = {
  id : '0',
  school: '',
  major : '',
  degree: '',
  start: null,
  end: null,
  isCurrent: false,
}

export default {
  name: 'Educations',

  data() {
    return {
      educations: {
        list: []
      },
      educationsRules: {
        school: [
          { required: true, message: 'Please input school', trigger: 'blur'},
          { min: 2, max: 32, message: 'Text length must between 2 and 32', trigger: 'blur'}
        ],
        major: [
          { required: true, message: 'Please input major', trigger: 'blur'},
          { min: 2, max: 32, message: 'Text length must between 2 and 32', trigger: 'blur'}
        ],
        degree: [
          { required: true, message: 'Please input degree', trigger: 'blur'},
          { min: 2, max: 32, message: 'Text length must between 2 and 32', trigger: 'blur'}
        ],
        start: [
          { required: true, message: 'Please input major', trigger: 'blur'},
          { type: 'date', message: 'Start time must be a date type', trigger: ['blur', 'change'] },
        ],
        end: [
          { required: true, message: 'Please input major', trigger: 'blur'},
          { type: 'date', message: 'End time must be a date type', trigger: ['blur', 'change'] },
        ],
      }
    }
  },

  methods: {
    add() {
      this.educations.list.push(Object.assign({},education))
      console.log(this.educations)
    },
    del(i, id) {
      this.$confirm('Delete this education?', 'Alert', {
        confirmButtonText: 'Delete',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(() => {
        this.DEL_EDUCATION({
          'id': id.toString(),
          'user_id': this.$store.state.user_id.toString(),
        }).then((res) => {
          this.educations.list.splice(i, 1)
          this.$message({
            type: 'success',
            message: 'delete success!'
          });
        }, (error) => {
          this.$notify.error({
            title: 'Delete education failed.',
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
      this.EDUCATION({
        'id': this.$store.state.user_id.toString(),
        'list': this.educations.list}).then((res) => {
        this.$notify({
          title: 'Success',
          message: 'Edit educations success!',
          position: 'bottom-right',
          type: 'success'
        })
      }, (error) => {
        this.$notify.error({
          title: 'Edit educations failed.',
          message: error.response.data.msg,
          position: "top-right"
        })
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    ...mapActions([
      'EDUCATION',
      'GET_EDUCATION',
      'DEL_EDUCATION',
    ])
  },

  mounted() {
    this.GET_EDUCATION(this.$store.state.user_id.toString()).then((res) => {
      console.log('Get educations success!')
      if (res.data.educations.length > 0) {
        var educationList = []
        res.data.educations.forEach(function(eachEducation){
          var e = Object.assign({}, education)
          e.id = eachEducation.ID.toString()
          e.school = eachEducation.School
          e.major = eachEducation.Major
          e.degree = eachEducation.Degree
          e.start = eachEducation.StartTime
          e.end = eachEducation.EndTime
          e.isCurrent = eachEducation.IsCurrent
          educationList.push(e)
        })
        this.educations.list = educationList
      } else {
        this.educations.list = [Object.assign({},education)]
      }
    }, (error) => {
      this.$notify.error({
        title: 'Get educations failed.',
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

.education-card {
  text-align: left;
  margin-bottom: 30px;
}
</style>
