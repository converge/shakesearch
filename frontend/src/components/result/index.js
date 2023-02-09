import './style.css'

export const ResultItem = (props) => {
    return (
        <div className="query-result">
            <div className="query-result__box">
                <h3>title: {props.title}</h3>
                <pre>{props.content}</pre>...
            </div>
            <div className="query-result__more">
                <a href={"/chapter/" + props.id}>read more</a>
            </div>
        </div>
    );
};

export const ResultsList = (props) => {
    return (
        <div>
            {props.queryResult.map((result) => (
                <ResultItem key={result.id} id={result.id} title={result.title} content={result.content}/>
            ))}
        </div>
    );
};
