import {UserProfile} from "../types/User";
import {AllowanceType, EmployeeAllowance} from "../types/Allowance";

const API_BASE_URL = import.meta.env.VITE_API_URL as string;

export async function fetchAPI<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
    const response = await fetch(`${API_BASE_URL}${endpoint}`, {
        ...options,
        credentials: "include",
    });
    if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
    }
    return await response.json() as Promise<T>;
}

export async function fetchEmployees<T>(): Promise<T> {
    return fetchAPI<T>("/employees");
}

export async function fetchEmployeeById<T>(id: string): Promise<T> {
    return fetchAPI<T>(`/employees/${id}`);
}

export async function createAttendance<T>(
    employeeId: string,
    data: { check_in: string; check_out: string }
): Promise<T> {
    return fetchAPI<T>(`/employees/${employeeId}/attendances`, {
        method: 'POST',
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(data),
    });
}

export async function fetchPayroll<T>(employeeId: string, year: number, month: number): Promise<T> {
    return fetchAPI<T>(`/employees/${employeeId}/payroll?year=${year}&month=${month}`);
}

export async function login<T>(credential: { email: string, password: string }): Promise<T> {
    return fetchAPI<T>("/login", {
        method: 'POST',
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(credential),
    });
}

export async function logout<T>(): Promise<T> {
    return fetchAPI<T>("/logout", {method: "POST"});
}

export async function fetchProfile<T = UserProfile>(): Promise<T> {
    return fetchAPI<T>("/current_account");
}

export async function fetchAllowanceTypes<T>(): Promise<T> {
    return fetchAPI<T>("/allowance_types");
}

export async function createAllowanceType<T>(data: Partial<AllowanceType>): Promise<T> {
    return fetchAPI<T>("/allowance_types", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(data),
    });
}

export async function updateAllowanceType<T>(id: number, data: Partial<AllowanceType>): Promise<T> {
    return fetchAPI<T>(`/allowance_types/${id}`, {
        method: "PUT",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(data),
    });
}

export async function deleteAllowanceType<T>(id: number): Promise<T> {
    return fetchAPI<T>(`/allowance_types/${id}`, {
        method: "DELETE",
    });
}

export async function fetchEmployeeAllowances<T>(): Promise<T> {
    return fetchAPI<T>("/employee_allowances");
}

export async function createEmployeeAllowance<T>(data: EmployeeAllowance): Promise<T> {
    return fetchAPI<T>("/employee_allowances", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(data),
    });
}

export async function updateEmployeeAllowance<T>(id: number, data: EmployeeAllowance): Promise<T> {
    return fetchAPI<T>(`/employee_allowances/${id}`, {
        method: "PUT",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(data),
    });
}

export async function deleteEmployeeAllowance<T>(id: number): Promise<T> {
    return fetchAPI<T>(`/employee_allowances/${id}`, {
        method: "DELETE",
    });
}
