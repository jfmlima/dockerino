import React, {useEffect, useState} from "react";
import axios, {AxiosResponse} from "axios";

interface Response {
    value: string;
}

const Docker: React.FC = () => {
    const [text, setText] = useState("");

    useEffect(() => {
        axios.get("api/whoami")
            .then((t: AxiosResponse<Response>) => setText(t.data.value))
    }, []);

    return <div>{text}</div>
};

export default Docker;
