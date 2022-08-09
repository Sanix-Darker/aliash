import { Link } from 'preact-router/match';
import {useState} from "preact/hooks";
import {dataSearchAliases} from "../../data";
import './style.css';


const SearchItem = ({key, item}) => {

    return <Link href={`/details/${item.Uid}`}>
        <li key={key} className="item">{item?.Title}</li>
    </Link>
}

const SearchListItem = ({searchList, query}) => {

    return searchList.length > 0 && query.length > 0 ? (
            <div className="result-block">
                <ul> {searchList.map((item, key) => <SearchItem key={key} item={item} />)} </ul>
            </div>
    ): null
}

const SearchBox = () => {

    const [query, setQuery] = useState('');
    const [searchList, setSearchList] = useState([])

    const handleAutoCompletion = async (searchValue) => {
        setQuery(searchValue);
        if (searchValue.length > 1){
            // added setimeout to split amout of backend calls
            setTimeout(async () => {
                const res = await dataSearchAliases(searchValue);
                setSearchList(res);
            }, 500);
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
        <SearchListItem searchList={searchList} query={query} />
    </>;
}

export default SearchBox;
