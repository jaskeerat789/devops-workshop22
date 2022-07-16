const express = require('express')
const mongoose = require('mongoose')
require('dotenv').config()
const app = express()
const port = process.env.PORT || 3000

app.get('/', (req, res) => {
    res.send("<h1>Welcome to DevOps Workshop</h1><p>Node.js app v2.0.</p>")
})


mongoose.connect('mongodb://mongo:27017/workshop', {
    useNewUrlParser: true,
})
.then(() => {
    console.log("Connected to the DB")
    app.listen(port, () => {
        console.log(`Blazing fast server in running on port ${port}`)
    })
})
.catch(err => {
    console.log("failed to connect to mongodb",err)
})
