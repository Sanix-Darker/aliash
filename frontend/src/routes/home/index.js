import SearchBox from '../../components/searchBox';
import style from './style.css';

const Home = () =>{

    return (
        <div class={style.home}>
            <h1>AliasHub</h1>
            <p>
                Search, Host and Use shell aliases related to
                installations/settings or configuration on unix OS.
            </p>
            <SearchBox />
        </div>
    );
}

export default Home;
