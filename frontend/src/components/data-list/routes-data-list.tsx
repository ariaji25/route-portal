import { Button, CircularProgress, Paper, Stack, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Typography, Alert } from "@mui/material";
import { useState } from "react";
import RouteFormModal from "../modal/route-form-modal";
import ConfirmationModal from "../modal/confirmation-modal";
import { Route } from "@/types/models/route";
import { useGetRoutesQuery, useDeleteRouteMutation } from "@/services/route-service";

export default function RoutesDataList() {
    const [openFormModal, setOpenFormModal] = useState(false);
    const [editingRoute, setEditingRoute] = useState<Route | undefined>(undefined);
    const [deleteConfirmOpen, setDeleteConfirmOpen] = useState(false);
    const [routeToDelete, setRouteToDelete] = useState<string>("");
    
    const { data: routes, isLoading, error } = useGetRoutesQuery();
    const deleteRouteMutation = useDeleteRouteMutation();

    const handleAddRoute = () => {
        setEditingRoute(undefined);
        setOpenFormModal(true);
    };

    const handleEditRoute = (route: Route) => {
        setEditingRoute(route);
        setOpenFormModal(true);
    };

    const handleDeleteClick = (routeName: string) => {
        setRouteToDelete(routeName);
        setDeleteConfirmOpen(true);
    };

    const handleDeleteConfirm = async () => {
        if (routeToDelete) {
            try {
                await deleteRouteMutation.mutateAsync(routeToDelete);
                setDeleteConfirmOpen(false);
                setRouteToDelete("");
            } catch (error) {
                console.error('Failed to delete route:', error);
            }
        }
    };

    const handleDeleteCancel = () => {
        setDeleteConfirmOpen(false);
        setRouteToDelete("");
    };

    const handleFormModalClose = () => {
        setOpenFormModal(false);
        setEditingRoute(undefined);
    };

    return (
        <Stack p={6} gap={2}>
            <Stack direction={"row"} justifyContent={"center"} alignItems={"center"}>
                <Typography variant="h6" gutterBottom>
                    Routes Data List
                </Typography>
                <Button onClick={handleAddRoute} variant="contained" color="primary" sx={{ marginLeft: 'auto' }}>
                    Add Route
                </Button>
            </Stack>
            
            {error && (
                <Alert severity="error">
                    Failed to load routes. Please try again.
                </Alert>
            )}
            {
                isLoading
                    ? (<CircularProgress />)
                    : <TableContainer component={Paper}>
                        <Table>
                            <TableHead>
                                <TableRow>
                                    <TableCell>Name</TableCell>
                                    <TableCell>Host</TableCell>
                                    <TableCell>Backend</TableCell>
                                    <TableCell>Path</TableCell>
                                    <TableCell>Enabled</TableCell>
                                    <TableCell>Actions</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {routes?.data?.map((route) => (
                                    <TableRow key={route.name}>
                                        <TableCell>{route.name}</TableCell>
                                        <TableCell>{route.host}</TableCell>
                                        <TableCell>{route.backend}</TableCell>
                                        <TableCell>{route.path}</TableCell>
                                        <TableCell>{route.enabled ? "Yes" : "No"}</TableCell>
                                        <TableCell>
                                            <Button 
                                                variant="outlined" 
                                                size="small" 
                                                color="info"
                                                onClick={() => handleEditRoute(route)}
                                                disabled={deleteRouteMutation.isPending}
                                            >
                                                Edit
                                            </Button>
                                            <Button 
                                                variant="outlined" 
                                                size="small" 
                                                color="error" 
                                                sx={{ marginLeft: 1 }}
                                                onClick={() => handleDeleteClick(route.name)}
                                                disabled={deleteRouteMutation.isPending}
                                            >
                                                Delete
                                            </Button>
                                        </TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                        </Table>
                    </TableContainer>
            }
            
            <RouteFormModal 
                open={openFormModal} 
                onClose={handleFormModalClose}
                editingRoute={editingRoute}
            />
            
            <ConfirmationModal
                open={deleteConfirmOpen}
                onClose={handleDeleteCancel}
                onConfirm={handleDeleteConfirm}
                title="Confirm Delete"
                message={`Are you sure you want to delete the route "${routeToDelete}"? This action cannot be undone.`}
                confirmText="Delete"
                cancelText="Cancel"
                confirmColor="error"
                isLoading={deleteRouteMutation.isPending}
            />
        </Stack>
    )
}