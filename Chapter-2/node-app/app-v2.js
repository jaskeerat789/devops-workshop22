const express = require('express')
require('dotenv').config()
const app = express()
const port = process.env.PORT ||  3000

app.get('/',(req,res) => {
    res.send("<h1>Welcome to DevOps Workshop</h1><p>Node.js app v2.0.</p>")
})

app.listen(port,()=>{
    console.log(`Blazing fast server in running on port ${port}`)
})