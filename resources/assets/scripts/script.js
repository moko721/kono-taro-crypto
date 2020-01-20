const url = location.href;

const onButtonClick = () => {
    fetch(`${url}api/new`)
        .then(res => res.json())
        .then(json => {
            document.getElementById('hashedText').innerHTML = json.hashed;
        })
        .catch(err => {
            console.error(err)
        })
};