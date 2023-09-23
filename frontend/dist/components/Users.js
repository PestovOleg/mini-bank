var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
export function Users() {
    let content = '<div class="users bg-info-subtle">';
    content += '<h2>Create users 111</h2>';
    content += '<h2>Delete Users</h2>';
    content += '<h2>Update Users</h2>';
    content += '<h2>Get User</h2>';
    content += '</div>';
    return content;
}
function getStatus() {
    return __awaiter(this, void 0, void 0, function* () {
        const response = yield fetch('/api/v1/health');
        if (!response.ok) {
            throw new Error('HTTP error! Cannot get server status');
        }
        return yield response.json();
    });
}
