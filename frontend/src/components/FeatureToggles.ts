export function FeatureToggles(){
    const features=['CreateUserToggle','UpdateUserToggle','InfoUserToggle','DeleteUserToggle','CreateAccountToggle'];
    let content ='<div class="features-toggles">'
    features.forEach(feature=>{
        content+='<label><input type="checkbox">${feature}</label><br>'
    });
    content+='</div>';
    return content;
}