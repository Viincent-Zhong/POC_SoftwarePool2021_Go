import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import "./components/message/index.css"
import Message from "./components/message"
import My_login from "./components/message/login"

ReactDOM.render(
  <React.StrictMode></React.StrictMode>,
  document.getElementById("root")
);

ReactDOM.render(
  <React.StrictMode>
    <Message
      messageId={1}
      imgProfileUrl="https://i.ibb.co/Npv3bk7/image-2021-02-18-164910.png"
      senderName="Pepe"
      messageText="Life sucks, is banana a bread, gonna kill a ship later"
    />,
    <My_login
    username="jeanluc"
    password="envie de die"
    profilePicture="https://i.ibb.co/Npv3bk7/image-2021-02-18-164910.png"
    />,
  </React.StrictMode>,
  document.getElementById("root")
);

/*
ReactDOM.render(
  <React.StrictMode>
    <My_login
    username="jeanluc"
    password="envie de die"
    profilePicture="https://i.ibb.co/Npv3bk7/image-2021-02-18-164910.png"
    />,
  </React.StrictMode>,
  document.getElementById("root")
);*/

/*
class My_form extends React.Component {
  render() {
    return(
      <form>
        <p>Send a message: </p>
        <input type="text"/>
      </form>
    );
  }
}
ReactDOM.render(<My_form/>, document.getElementById('root'))*/