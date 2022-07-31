function downloadFile(){
    let key = window.location.href.split("/").pop();
    
    window.open("/api/dl/"+key)
}