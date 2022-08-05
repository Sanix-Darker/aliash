const HOST = process.env.HOST
export const searchAliases = async (searchText) => {
    const res = await fetch(`${HOST}/search?q=${searchText}`);
    const result = res.json();
    return result;
}
