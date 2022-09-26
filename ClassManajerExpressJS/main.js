const express = require('express')
const app = express()
const port = 3000

// HOME
app.get('/', function(req, res){
  res.send({"message": "index"})
})

// UNSECURE ROUTE
app.use('/', require('./controllers/all_user'))
app.use('/', require('./controllers/login'))
app.use('/', require('./controllers/register'))


// START PROGRAM
app.listen(port, () => {
  console.log(`listening at http://localhost:${port}`)
});
