// init express.js
const express = require('express')

// init db
const { initializeApp } = require('firebase/app')
const { getDatabase, ref , get, set } = require('firebase/database')

const firebase_config = require('../firebaseConfig/config')

initializeApp(firebase_config)
const db = getDatabase();

// main func
var AllUsers = express.Router()

AllUsers.get('/allstudents', function(req, res){
    get(ref(db, 'user/student')).then((snap) => {
        if (snap.exists()){
            res.send({
              "student" : snap
          })
        } else {
            res.send({
              "message" : "error"
            })
        }
    })
    console.log()
})

AllUsers.get('/allteachers', function(req, res){
    get(ref(db, 'user/teacher')).then((snap) => {
        if (snap.exists()){
            res.send({
              "teacher" : snap
          })
        } else {
            res.send({
              "message" : "error"
            })
        }
    })
    console.log()
})

module.exports = AllUsers
