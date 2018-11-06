module.exports = {
    "env": {
        "browser": true,
        "es6": true
    },
    "extends": "plugin:vue/essential",
    "parserOptions": {
        "parser": "babel-eslint",
        "ecmaVersion": 2015,
        "sourceType": "module"
    },
    "rules": {
        "indent": [
            "error",
            2
        ],
        "linebreak-style": [
            "error",
            "unix"
        ],
        "semi": 0
    }
};