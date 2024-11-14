import express from "express";
import { clog, cerr } from "easier-jsever";
import pool from "../configuration/databasecon.js";

const router = express.Router();

router.get("/", async (req, res) => {
    try {
        const userQuery = "SELECT COUNT(DISTINCT email) AS email_count FROM techsurvey.users";
        const [result] = await pool.query(userQuery);

        if (result.length === 0) {
            console.log("No number of emails found");
            return res.status(404).json({ message: "No number of emails found" });
        }

        // Access the count from the result
        const emailCount = result[0].email_count;

        // Log the result (for debugging purposes)
        clog(result);

        // Return the email count in a proper JSON format
        res.status(200).json({email_count: emailCount });
    } catch (error) {
        // Log detailed error and send server error response
        cerr("Database query error:", error);
        res.status(500).json({ error: "Internal server error" });
    }
});

export default router;
