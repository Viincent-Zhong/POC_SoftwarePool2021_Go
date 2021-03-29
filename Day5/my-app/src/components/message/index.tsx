import "./index.css";

type MessageProps = {
    messageId: number;
    imgProfileUrl: string;
    senderName: string;
    messageText: string;
}

function Message(props : MessageProps): JSX.Element {
    return (<div className = "text">
        <p>{props.messageId}
        <br></br>{props.senderName}
        <br></br>{props.messageText}</p>
        <img src = {props.imgProfileUrl}/>
        </div>
    )
}

export default Message