import { Button } from '@material-ui/core';


function handleClick() {
    url = "http://localhost:8082"
    fetch("")
    console.log("Hello");
    alert(1);
}

function Body(props) {
    return (
        <main className={props.main}>
        <h1 className={props.title}>
          Welcome Voz reader
        </h1>
        <Button onClick={handleClick} color="primary">Click me</Button>
      </main>
    )
}

export default Body