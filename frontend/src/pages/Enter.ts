import { Footer } from "../components/Footer.js";
import { Navbar } from "../components/Navbar.js";
import { SignIn } from "../components/SignIn.js";
import { SignUp } from "../components/SignUp.js";
import { Component} from "../components/component.js"

export function Enter(): string {
    let signInContent = SignIn();
    let signUpContent = SignUp();
    let footerContent = Footer();
    let navbarContent=Navbar();
    let content= `
    <div class="container-lg bg-primary-subtle
     enter-page d-flex flex-column justify-content-between 
     align-items-center vh-100">
        ${navbarContent}     
            <div class="second-div-enter-page d-flex flex-column align-items-center justify-content-center">
            ${signInContent}
            ${signUpContent}
            </div>
        ${footerContent}
    </div>`;
    content+=Component()
    return content
}