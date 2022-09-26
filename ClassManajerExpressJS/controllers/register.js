// init express.js
const express = require('express')

// init db
const { initializeApp } = require('firebase/app')
const { getDatabase, ref , get, set } = require('firebase/database')

const firebase_config = require('../firebaseConfig/config')

initializeApp(firebase_config)
const db = getDatabase();

// main func
var Register = express.Router()

Register.post('/registerstudent', (req, res) => {
    var i = 0
    get(ref(db, 'user/student')).then((snap) => {
      var users = snap.val()
      
      for (const user in users){
        i++
      }
      i++
      
      var adduser = set(ref(db, 'user/student/' + i.toString()), req.body)
  
      if (adduser){
        res.send({
          "status":"success",
          "content" : req.body
        })
      } else {
        res.send({
          "status":"gagal",
          "content" : req.body
        })
      }
    })
  })
  
  Register.post('/registerteacher', (req, res) => {
    
    var i = 0
    get(ref(db, 'user/teacher')).then((snap) => {
      var users = snap.val()
      
      for (const user in users){
        i++
      }
      i++
      
      var adduser = set(ref(db, 'user/teacher/' + i.toString()), req.body)
  
      if (adduser){
        res.send({
          "status":"success",
          "content" : req.body
        })
      } else {
        res.send({
          "status":"gagal",
          "content" : req.body
        })
      }
    })
  })

  module.exports = Register
