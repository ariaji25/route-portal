import { Modal, Stack } from "@mui/material";
import RouteForm from "../form/route-form";
import { Route } from "@/types/models/route";
import { useCreateRouteMutation, useUpdateRouteMutation } from "@/services/route-service";

interface RouteFormModalProps {
    open: boolean;
    onClose: () => void;
    editingRoute?: Route;
}

export default function RouteFormModal({
    open,
    onClose,
    editingRoute
}: RouteFormModalProps) {
    const createRouteMutation = useCreateRouteMutation();
    const updateRouteMutation = useUpdateRouteMutation();

    const handleSubmit = async (data: Route) => {
        try {
            if (editingRoute) {
                await updateRouteMutation.mutateAsync({ name: editingRoute.name, data });
            } else {
                await createRouteMutation.mutateAsync(data);
            }
            onClose();
        } catch (error) {
            console.error('Failed to save route:', error);
            throw error; // Re-throw to let the form handle the error state
        }
    };

    const isLoading = createRouteMutation.isPending || updateRouteMutation.isPending;

    return (
        <Modal open={open} onClose={onClose}>
            <Stack 
                direction={"column"} 
                justifyContent={"center"} 
                alignItems={"center"} 
                sx={{ 
                    height: '100vh',
                    overflow: 'auto',
                    p: 2
                }}
            >
                <RouteForm 
                    onSubmit={handleSubmit}
                    initialData={editingRoute}
                    isLoading={isLoading}
                    onCancel={onClose}
                />
            </Stack>
        </Modal>
    )
}