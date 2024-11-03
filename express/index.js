import express from "express";

import cors from "cors";

import fetchdata from "./Routes/getSurveyData.js";

const port = 6900;

import path from "path";

import helmet from "helmet";

import dotenv from "dotenv";

dotenv.config();

import {clog,cerr} from "easier-jsever";

const app = express();

app.use(express.urlencoded({ extended: false }));

app.use(express.json());

app.use(helmet());

app.use("/getdata",fetchdata);

app.use(cors("http://localhost:4700"))

app.listen(port , ()=>clog(`server is running on port ${port}`));