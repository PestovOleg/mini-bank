import styles from "../types/styles.module.css"

export function Navbar(): string {
    let content = `
        <nav class="navbar-expand-lg navbar-div container-lg">
            <div class="container-fluid ">
                <p class=`+`${styles.minibankBrand}`+`text-center>MINIBANK</p>
            </div>            
        </nav>`
    return content
}