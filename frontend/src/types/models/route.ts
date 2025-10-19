export interface Response<T> {
    data: T;
    message: string;
    status: number;
}

export interface Route {
    name: string;
    host: string;
    backend: string;
    path: string;
    enabled: boolean;
}   