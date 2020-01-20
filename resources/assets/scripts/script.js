const url = location.href;
const select = document.getElementById('cost');

for (let i = 4; i <= 32; i++) {
    let option = document.createElement('option');
    let v = 1;
    for (let j = 1; j <= i; j++) {
        v *= 2;
    }
    option.value = i.toString();
    option.innerHTML = v.toString() + `(2^${i})`;
    select.appendChild(option);
}

const onButtonClick = () => {
    const cost = document.getElementById('cost').value;
    fetch(`${url}api/new?cost=${cost}`)
        .then(res => res.json())
        .then(json => {
            document.getElementById('hashedText').innerHTML = json.hashed;
        })
        .catch(err => {
            console.log(err)
        })
};