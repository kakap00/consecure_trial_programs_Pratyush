function ListGroup() {
  const items = ["bangalore", "mumbai", "chennai", "assam", "jodhpur"];

  return (
    //we cannot create a header h1 along with this ul list as a singular component can only
    //return 1 singular thing. hence, we can wrap both h1 and ul in another empty <>
    //empty <> is basically a shorthand notation for fragmentation. (import fragment and instead)
    //(of empty <> we'd have had to put <fragment>)

    //note that items.map() must be enclosed in curly braces inside the ul tag as it only allows html
    <>
      <h1>LIST ITEMS</h1>
      <ul className="list-group"> 
        {items.map((item) => (
          <li className="list-group-item" 
          key = {item} 
          onClick = {() => console.log("clicked")}>
            {item}
          </li>
        ))}
      </ul>
    </>
  );
}

export default ListGroup;
