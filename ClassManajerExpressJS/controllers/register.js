// init express.js
const express = require('express')

// init db
const { initializeApp } = require('firebase/app')
const { getDatabase, ref , set, get } = require('firebase/database')

const firebase_config = require('../firebaseConfig/config')

initializeApp(firebase_config)
const db = getDatabase();

// include functions
const { randomStr } = require('../functions/random_str')

// main
var Register = express.Router()

Register.post('/register', async (req, res) => {

  // get data
  const username  = req.body.username
  const email     = req.body.email
  const password  = req.body.password
  const type      = req.body.type

  // check payload
  if (!username || !email || !password || !type){

    res.status(400).send({
      'message' : 'null payload'
    })

    return 
  } 

  // check duplicate email
  duplicated_email = false
  await get(ref(db, 'user/' + type.toString() + '/')).then((snap) => {

    emailList = snap.val()

    for (const address in emailList){
      
      if (eval('emailList.' + address + '.email').toString() == email.toString()){
        duplicated_email = true
      }
    }
  })

  if (duplicated_email == true){

    res.send({
      "message"  :"duplicate email"
    })

    return
  }

  // construct json
  const value = {
    "username"  : username,
    "password"  : password,
    "email"     : email
  }

  const register = set(ref(db, 'user/' + type + '/' + randomStr(30)), value)
  
  if (register){
    res.send({
      "status":"success",
      "content" : value
    })
  } else {
    res.send({
      "status":"failed to register",
      "content" : value
    })

    return
  }
})

module.exports = Register
