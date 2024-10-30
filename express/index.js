const express = require("express");
const app = express();
const cors = require("cors");
app.use(cors())


app.listen(4200,()=>{
    console.log("server started at 4200")
})