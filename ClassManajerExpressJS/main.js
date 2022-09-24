const { initializeApp } = require('firebase/app');
const { getDatabase, ref , get, set, goOffline } = require('firebase/database');

const express = require('express')
const bodyParser = require('body-parser')
const app = express();
const port = 3000;

const firebaseConfig = {
    apiKey: "AIzaSyCAV4IZbWdnBo__3MQFl9nfHY016sqP6hY",
    authDomain: "classmanagerapps.firebaseapp.com",
    projectId: "classmanagerapps",
    storageBucket: "classmanagerapps.appspot.com",
    messagingSenderId: "1023940208070",
    appId: ""
  };


initializeApp(firebaseConfig)
const db = getDatabase();

app.use(bodyParser.json());

app.get('/uname', (req, res) => {

  var username;
  get(ref(db, 'user/student')).then((snap) => {
      if (snap.exists()){
          username = snap.val()
      }
      res.send(username)
  });
});

app.post('/login', (req, res) => {
  var body = req.body
  
  res.send(body)

});

app.post('/adduser', (req, res) => {
  var adduser = set(ref(db, 'new/val'), req.body)

  if (adduser){
    res.send({
      "status":"success",
      "content": req.body
    })
  } else {
    res.send("gagal")
  }

  
})

app.listen(port, () => {
  console.log(`listening at http://localhost:${port}`)
});





 
