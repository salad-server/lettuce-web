import Swal from "sweetalert2";

export const API = () => {
    Swal.fire({
        title: "API Error!",
        text: "Check your connection. Please report this to a staff member if the problem persists.",
        icon: "error",
    });
};
