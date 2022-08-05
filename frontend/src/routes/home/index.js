import {useState} from 'preact/hooks';
import style from './style.css';

const Home = () =>{

    const [searchText, setSearchText] = useState('')

    return (
        <div class={style.home}>
            <h1>AliasHub</h1>
            <p>
                Search/Host shell aliases related to installations/settings for any unix OS
            </p>
            <input type="text"
                   value={searchText}
                   onChange={(e) => setSearchText(e.target.value)}
            />
        </div>
    );
}

export default Home;
