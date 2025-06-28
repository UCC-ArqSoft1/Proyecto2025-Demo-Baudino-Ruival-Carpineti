export function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
  }
  
  export function decodeJWT(token) {
    try {
      const base64Url = token.split('.')[1];
      const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
      const payload = JSON.parse(atob(base64));
      return payload;
    } catch (error) {
      console.error("Error decoding JWT:", error);
      return null;
    }
  }

export function getUserRole() {
  const token = getCookie("token");
  if (!token) return null;
  const payload = decodeJWT(token);
  return payload?.rol || payload?.role || null;
}