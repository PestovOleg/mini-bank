export function Users():string {
    let content = '<div class="users bg-info-subtle">';
    content += '<h2>Create users 111</h2>';
    content += '<h2>Delete Users</h2>';
    content += '<h2>Update Users</h2>';
    content += '<h2>Get User</h2>';
    content+='</div>'
    return content;
}

async function getStatus(){
    const response= await fetch('/api/v1/health');
    if (!response.ok){
        throw new Error('HTTP error! Cannot get server status')
    }
    return await response.json();
}