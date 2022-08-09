import env from '../../conf'

export const dataSearchAliases = async (searchText) => {
    const res = await fetch(`${env.API_URL}/search?q=${searchText}`);
    return await res.json();
}
