import styles from "../types/styles.module.css"

export function Footer():string{
    let content=`
    <div class="footer d-flex flex-column justify-content-between align-items-center container-lg">
        <a href="https://github.com/PestovOleg/mini-bank">Github</a>
        <span class="{styles.minibankBrand}" styles="background-color:black;">&copy; 2023 Pestov Oleg</span>
    </div>`
    return content
}