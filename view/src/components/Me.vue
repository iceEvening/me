<template>
  <div class="me">
    <el-row style="padding: 40px 0;">
      <el-col :span="22" :offset="1">
        <el-row>
          <el-col :span="7">
            <el-row>
              <el-col :span="24">
                <img :src="profile.img" style="margin-bottom: 20px">
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="24" style="word-wrap:break-word;">
                Hello, I'm <b>{{ profile.name }}</b>
                <i v-if="profile.gender === '1'" class="el-icon-ali-male_icon" style="margin-left:3px;"></i>
                <i v-if="profile.gender === '2'" class="el-icon-ali-female_icon" style="margin-left:3px;"></i>.<br>
                I was born in {{ profile.hometown}}<br>
                on {{ profile.birthday }}.<br>
                I'm working and living in <b>{{ profile.city }}</b> now.<br>
                You can contact me through this email: <br>
                <i class="el-icon-message"></i> {{ profile.email }}
              </el-col>
            </el-row>
          </el-col>
          <el-col :span="17" style="text-align: left; padding:0 10px;">
            <el-row style="height:150px;">
              <el-col :span="24">
                <b>Me:</b><br>
                <div v-html="profile.me" style="border-bottom: 1px solid #e6e6e6; margin-bottom:30px; padding-bottom: 20px;">
              </el-col>
            </el-row>
            <el-row style="margin:30px 0;">
              <el-col :span="24">
                Here is my work experience: <br>
                <el-table
                  :data="careers.list"
                  style="width: 100%">
                  <el-table-column
                    prop="company"
                    label="Company">
                  </el-table-column>
                  <el-table-column
                    prop="department"
                    label="Department"
                    min-width="85px">
                  </el-table-column>
                  <el-table-column
                    prop="post"
                    label="Post"
                    min-width="105px">
                  </el-table-column>
                  <el-table-column
                    prop="description"
                    label="Desc"
                    min-width="300px">
                  </el-table-column>
                  <el-table-column
                    prop="start"
                    label="Start time">
                  </el-table-column>
                  <el-table-column
                    prop="end"
                    label="End time">
                  </el-table-column>
                </el-table>
              </el-col>
            </el-row>
            <el-row style="margin:30px 0;">
              <el-col :span="24">
                Here is my education experience: <br>
                <el-table
                  :data="educations.list"
                  style="width: 100%">
                  <el-table-column
                    prop="school"
                    label="School">
                  </el-table-column>
                  <el-table-column
                    prop="major"
                    label="Major">
                  </el-table-column>
                  <el-table-column
                    prop="degree"
                    label="Degree">
                  </el-table-column>
                  <el-table-column
                    prop="start"
                    label="Start time"
                    width="75px">
                  </el-table-column>
                  <el-table-column
                    prop="end"
                    label="End time"
                    width="75px">
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
import { mapMutations } from 'vuex'

var career = {
  id : '0',
  company: '',
  department : '',
  post: '',
  rank: '',
  description: '',
  start: null,
  end: null,
  isCurrent: false,
}

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
  name: 'Me',
  
  data: function() {
    return {
      profile: {
        img: null,
        nickname: '',
        me: '',
        name: '',
        gender: '',
        email: '',
        birthday: null,
        hometown: '',
        city: '',
      },

      careers: {
        list: [],
      },

      educations: {
        list: [],
      },
    }
  },

  methods: {
    ...mapActions([
      'GET_USER_PROFILE',
      'GET_CAREERS',
      'GET_EDUCATIONS',
      'GET_EDUCATION',
      'GET_USER_PROFILE_IMAGE',
    ]),
    ...mapMutations([
      'UPDATE_PROFILE_TIME',
      'UPDATE_OWNIMG_IMG',
    ])
  },

  mounted() {
    this.GET_USER_PROFILE(this.$store.state.ownerID).then((res) => {
      this.profile.nickname = res.data.nickname
      this.profile.me = res.data.me
      this.profile.name = res.data.name
      this.profile.gender = res.data.gender.toString()
      this.profile.email = res.data.email
      this.profile.birthday = new Date(res.data.birthday).toDateString()
      this.profile.hometown = res.data.hometown
      this.profile.city = res.data.city
      if(this.$store.state.update != res.data.update) {
        this.UPDATE_PROFILE_TIME(res.data.update)
        this.GET_USER_PROFILE_IMAGE(this.$store.state.ownerID).then((res) => {
          this.UPDATE_OWNIMG_IMG(res.data.img)
          this.profile.img = 'data:image/jpeg;base64,' + res.data.img
        }, (error) => {
          this.$notify.error({
            title: 'Load user profile img failed.',
            message: error.response.data.msg,
            position: "top-right"
          })
        })
      } else {
        this.profile.img = 'data:image/jpeg;base64,' + this.$store.state.ownImg
      }
    }, (error) => {
      this.$notify.error({
        title: 'Load user profile failed.',
        message: error.response.data.msg,
        position: "top-right"
      })
    }),

    this.GET_CAREERS(this.$store.state.ownerID).then((res) => {
      console.log('Get careers success!')
      if (res.data.careers.length > 0) {
        var careerList = []
        res.data.careers.forEach(function(eachCareer){
          var c = Object.assign({}, career)
          c.id = eachCareer.ID
          c.company = eachCareer.Company
          c.department = eachCareer.Department
          c.post = eachCareer.Post
          c.rank = eachCareer.JobLevel
          c.description = eachCareer.JobDesc
          c.start = new Date(eachCareer.StartTime).getFullYear() + '/' + (new Date(eachCareer.StartTime).getMonth() + 1)
          c.end = new Date(eachCareer.EndTime).getFullYear() + '/' + (new Date(eachCareer.EndTime).getMonth() + 1)
          c.isCurrent = eachCareer.IsCurrent
          if (eachCareer.IsCurrent) {
            c.end = 'now'
          }
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
    }),

    this.GET_EDUCATION(this.$store.state.ownerID).then((res) => {
      console.log('Get educations success!')
      if (res.data.educations.length > 0) {
        var educationList = []
        res.data.educations.forEach(function(eachEducation){
          var e = Object.assign({}, education)
          e.id = eachEducation.ID
          e.school = eachEducation.School
          e.major = eachEducation.Major
          e.degree = eachEducation.Degree
          e.start = new Date(eachEducation.StartTime).getFullYear() + '/' + (new Date(eachEducation.StartTime).getMonth() + 1)
          e.end = new Date(eachEducation.EndTime).getFullYear() + '/' + (new Date(eachEducation.EndTime).getMonth() + 1)
          if (eachEducation.IsCurrent) {
            e.end = 'now'
          }
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
img {
  width: auto;
  height: auto;
  max-width: 100%;
  max-height: 100%;
}
</style>
