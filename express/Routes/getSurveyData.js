import express from 'express';
const router = express.Router();
import pool from "../configuration/databasecon.js";
import { clog ,cerr } from 'easier-jsever';
 router.get("/",async (req,res)=>{
    try{
     const [result] = await pool.query("select * from techsurvey.faveprogramlang")
     if (result.length === 0) {
        console.info('No survey data found');
        return res.status(404).json({ message: "No survey data found" });
    }
    clog(result);
    res.status(200).json(result);
    }
    catch (err) {
        // Log detailed error and send server error response
        cerr("Database query error:", err);
        res.status(500).json({ error: "Internal server error" });
    }
});
export default router;