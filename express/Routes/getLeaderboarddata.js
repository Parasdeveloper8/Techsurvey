import express from "express";
const router = express.Router();
import pool from "../configuration/databasecon.js";
import { cerr, cinfo ,clog} from "easier-jsever";
router.get("/",async(req,res)=>{
    try{
       const query = "select username,points from techsurvey.users order by points desc";
       const [result] = await pool.query(query);
       if (result.length === 0) {
        cinfo('No data found');
        return res.status(404).json({ message: "No data found" });
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