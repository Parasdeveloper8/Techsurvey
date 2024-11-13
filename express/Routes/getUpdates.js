import express from 'express';
const router = express.Router();
import pool from "../configuration/databasecon.js";
import { clog ,cerr, cinfo } from 'easier-jsever';
router.get("/",async (req,res)=>{
    try{
       const dataQuery ="select * from techsurvey.updates order by time desc";
       const[result] = await pool.query(dataQuery);
       if(result.length === 0){
        cinfo("no updates found")
           return res.json({message:"No data found"});
       }
       clog(result);
       res.status(200).json(result);
    }
    catch(err){
        // Log detailed error and send server error response
        cerr("Database query error:", err);
        res.status(500).json({ error: "Internal server error" });
    }
});
export default router;