import {useState} from "preact/hooks";
import {dataSearchAliases} from "../../data";
import './style.css';


const SearchBox = () => {
    const [query, setQuery] = useState('');
    const [searchList, setSearchList] = useState([])

    const handleAutoCompletion = async (searchValue) => {

        setQuery(searchValue);
        if (searchValue.length > 1){
            const res = await dataSearchAliases(searchValue);
            setSearchList(res);
        }
    }

    return <>
        <input
            type="text"
            placeholder="Type something..."
            autocomplete="off"
            autoFocus=''
            defaultValue={query}
            onKeyUp={(e) => handleAutoCompletion(e.target.value)}
        />
        {searchList.length > 0 && query.length > 0 ? (
            <div className="result-block">
                <ul> {searchList.map((r, i) => <li key={i} className="item">{r?.Title}</li>)} </ul>
            </div>
        ): null}
    </>;
}

export default SearchBox;
