import express from "express";
import pool from "../configuration/databasecon.js";
const router = express.Router();
import { clog ,cerr } from 'easier-jsever';
router.get("/",async (req,res)=>{
     try{
         const [result] = await pool.query("select * from techsurvey.feedback") 
         if (result.length === 0) {
          console.info('No comments found');
          return res.status(404).json({ message: "No comments found" });
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