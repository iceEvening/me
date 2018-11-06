exports.install = function (Vue, str) {
  Vue.prototype.checkSpecificKey = function (str) {
    var specialKey = "[@`~!#$^&*()=|{}':;',\\[\\].<>/?~！#￥……&*（）——|{}【】‘；：”“'。，、？]‘'"; 
    for (var i = 0; i < str.length; i++) {
      if (specialKey.indexOf(str.substr(i, 1)) != -1) {
        return false
      }
    }
    return true
  }
}