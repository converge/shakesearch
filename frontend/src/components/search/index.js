import './style.css';
import axios from "axios";
import {useState} from 'react';
import {ResultsList} from "../result";

const shakeSearch = async (query) => {

    const API_URL = process.env.API_URL
    if (API_URL === undefined) {
        console.log("API_URL is undefined")
        return null
    }
    try {
        const response = await axios.get(`${API_URL}/v1/chapter/search?query=${query}`);
        if (response.status === 200) {
            return response.data;
        }
        return null
    } catch (err) {
        console.log(err);
    }
    return null
}

export const Search = () => {

    const [query, setQuery] = useState('');
    const [queryResult, setQueryResult] = useState('');

    const handleSubmit = (ev) => {
        ev.preventDefault();

        const queryShakeSearch = async () => {
            const result = await shakeSearch(query);
            console.log(result)
            setQueryResult(result);
        }
        queryShakeSearch();
    }

    return <div id="search">
        <form onSubmit={handleSubmit}>
            <input
                type="text"
                id="query"
                name="search"
                value={query}
                onChange={event => setQuery(event.target.value)}
                required
            />
            <button type="submit">Search</button>
        </form>

        <div>
            {queryResult && <ResultsList queryResult={queryResult} />}

            {/*{queryResult && queryResult.map((result) => (*/}
            {/*    <div key={result.id}>{JSON.stringify(result.title)}</div>*/}
            {/*))}*/}
                {/*// <Result id={queryResult.id}/>*/}
        </div>
    </div>;
}
