import "./index.css";
import axios from "axios";
import { useState, useEffect } from 'react';
import React from 'react';

type LoginProps = {
    username: string;
    password: string;
    profilePicture: string;
}

function handle_login(username:string, password:string, pp_url:string) {
    axios.post('http://127.0.0.1:8080/api/register', {
        username: username,
        password: password,
        profilePicture: pp_url
    }).then ((response) => {
        console.log(response);
        return
    },
        (error) => {
            console.log(error)
            return
    });
    console.log(username, password, pp_url);
}

function My_login(props : LoginProps): JSX.Element {
    const[username, setUsername] = useState("")
    const[password, setPassword] = useState("")
    const[pp_url, setPPurl] = useState("")
    return (<form>
        <input type="text" name="Username" placeholder="username" onChange={(e) => {
            setUsername(e.target.value)
        }}/>
        <input type="text" name="password" placeholder="password" onChange={(e) => {
            setPassword(e.target.value)}}/>
        <input type="text" name="pp URL" placeholder= "pp_url" onChange={(e) => {
            setPPurl(e.target.value)}}/>
        <button type="button" onClick={handle_login}>Register</button>
    </form>);
    /*axios.post('http://127.0.0.1:8000/api/register', {
        username: props.username,
        password: props.password,
        profilePicture: props.profilePicture
    }).then ((response) => {
        console.log(response);
        return response
    },
        (error) => {
            console.log(error)
            return error;
    });
    return (<div></div>)*/
}

export default My_login