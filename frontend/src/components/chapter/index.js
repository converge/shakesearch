import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import axios from "axios";

const findChapterById = async (id) => {

    const API_URL = process.env.API_URL

    if (API_URL === undefined) {
        console.log("API_URL is undefined")
        return null
    }

    try {
        const response = await axios.get(`${API_URL}/v1/chapter/${id}`);
        if (response.status === 200) {
            return response.data;
        }
        return null
    } catch (err) {
        console.log(err)
    }
    return 'Not Found'
}

export const Chapter = () => {

    const [content, setContent] = useState('');

    const { id } = useParams();

    useEffect(() => {

        const effect = async () => {
            const result = await findChapterById(id);
            console.log(result)
            setContent(result);
        }
        effect()
    }, []);


    return (
        <div className="chapter-content">
            {content === 'Not Found' ? 'chapter not found!' : '' }
            {content &&
                <div>
                <h1>{content.title}</h1>
                <pre>{content.content}</pre>
                </div>
            }
        </div>
    );
}
