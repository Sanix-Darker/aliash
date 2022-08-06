// const HOST = process.env.API_HOST_NAME
const HOST = "http://127.0.0.1:5000"
export const dataSearchAliases = async (searchText) => {
    const res = await fetch(`${HOST}/search?q=${searchText}`);
    return await res.json();
}
