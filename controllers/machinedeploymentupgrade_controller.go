/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// MachineDeploymentUpgradeReconciler reconciles a MachineDeploymentUpgrade object.
type MachineDeploymentUpgradeReconciler struct {
	client client.Client
}

// NewMachineDeploymentUpgradeReconciler returns a new instance of MachineDeploymentUpgradeReconciler.
func NewMachineDeploymentUpgradeReconciler(client client.Client) *MachineDeploymentUpgradeReconciler {
	return &MachineDeploymentUpgradeReconciler{
		client: client,
	}
}

//+kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=machinedeploymentupgrades,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=machinedeploymentupgrades/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=machinedeploymentupgrades/finalizers,verbs=update

// Reconcile reconciles a MachineDeploymentUpgrade object.
func (r *MachineDeploymentUpgradeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MachineDeploymentUpgradeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&anywherev1.MachineDeploymentUpgrade{}).
		Complete(r)
}
