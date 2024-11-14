import express from "express";

import cors from "cors";

import fetchfeed from "./Routes/getFeedback.js";

import fetchupdate from "./Routes/getUpdates.js";

import fetchUserNumber from "./Routes/getNumUser.js";
// Configure CORS
const corsOptions = {
    origin: 'http://localhost:4700', // Allow requests from your frontend's origin
    methods: ['GET', 'POST', 'PUT', 'DELETE', 'OPTIONS'], // Allowed methods
    allowedHeaders: ['Content-Type', 'Authorization'], // Allowed headers
    credentials: true // Allow credentials (if needed)
};
const app = express();

app.use(cors(corsOptions)); // Enable CORS with options

import fetchdata from "./Routes/getSurveyData.js";

const port = 7300;

import path from "path";

import helmet from "helmet";

import dotenv from "dotenv";

dotenv.config();

import {clog,cerr} from "easier-jsever";

app.use(express.urlencoded({ extended: false }));

app.use(express.json());

app.use(helmet());

app.use("/getdata",fetchdata);

app.use("/getfeedback",fetchfeed);

app.use("/getupdate",fetchupdate);

app.use("/getUserNumber",fetchUserNumber);

app.listen(port , ()=>clog(`server is running on port ${port}`));