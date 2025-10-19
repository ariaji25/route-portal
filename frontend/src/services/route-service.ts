import { Response, Route } from "@/types/models/route";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";

const baseApiUrl = 'http://localhost:8080';

export const routeServices = {
    async createRoute(data: Route): Promise<Response<Route>> {
        const response = await fetch(`${baseApiUrl}/routes/`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        });

        if (!response.ok) {
            throw new Error('Failed to create route');
        }

        return response.json();
    },

    async getRoutes(): Promise<Response<Route[]>> {
        const response = await fetch(`${baseApiUrl}/routes/`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error('Failed to fetch routes');
        }

        return response.json();
    },

    async updateRoute(name: string, data: Route): Promise<Response<Route>> {
        const response = await fetch(`${baseApiUrl}/routes/${name}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        });

        if (!response.ok) {
            throw new Error('Failed to update route');
        }

        return response.json();
    },

    async deleteRoute(name: string): Promise<void> {
        const response = await fetch(`${baseApiUrl}/routes/${name}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error('Failed to delete route');
        }
    }
};

const useGetRoutesQuery = () => {
    return useQuery<Response<Route[]>>({
        queryKey: ['routes'],
        queryFn: routeServices.getRoutes,
    });
}

const useCreateRouteMutation = () => {
    const queryClient = useQueryClient();
    return useMutation({
        mutationFn: routeServices.createRoute,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['routes'] });
        },
    });
}

const useUpdateRouteMutation = () => {
    const queryClient = useQueryClient();
    return useMutation({
        mutationFn: ({ name, data }: { name: string; data: Route }) => routeServices.updateRoute(name, data),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['routes'] });
        },
    });
}

const useDeleteRouteMutation = () => {
    const queryClient = useQueryClient();
    return useMutation({
        mutationFn: (name: string) => routeServices.deleteRoute(name),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['routes'] });
        },
    });
}

export {
    useGetRoutesQuery,
    useCreateRouteMutation,
    useUpdateRouteMutation,
    useDeleteRouteMutation
};
