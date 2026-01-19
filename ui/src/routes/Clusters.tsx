import { useEffect, useMemo } from "react";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination";
import ClusterCard from "@/components/ClusterCard";
import ClusterCompareDialog from "@/components/ClusterCompareDialog/ClusterCompareDialog";
import { useStore } from "@/store/useStore";

const ITEMS_PER_PAGE = 6;

const ClustersPagePagination = ({
  currentPage,
  totalPages,
  onPageChange,
}: {
  currentPage: number;
  totalPages: number;
  onPageChange: (page: number) => void;
}) => {
  const getPageNumbers = () => {
    const pages = [];
    const maxVisible = 5;
    let startPage = Math.max(1, currentPage - Math.floor(maxVisible / 2));
    let endPage = Math.min(totalPages, startPage + maxVisible - 1);

    if (endPage - startPage < maxVisible - 1) {
      startPage = Math.max(1, endPage - maxVisible + 1);
    }

    for (let i = startPage; i <= endPage; i++) {
      pages.push(i);
    }
    return pages;
  };

  return (
    <Pagination>
      <PaginationContent>
        <PaginationItem>
          <PaginationPrevious
            onClick={() => currentPage > 1 && onPageChange(currentPage - 1)}
            className={
              currentPage === 1 ? "pointer-events-none opacity-50 cursor-not-allowed" : "cursor-pointer"
            }
          />
        </PaginationItem>

        {getPageNumbers().map((pageNum) => (
          <PaginationItem key={pageNum}>
            <PaginationLink
              isActive={currentPage === pageNum}
              onClick={() => onPageChange(pageNum)}
              className="cursor-pointer"
            >
              {pageNum}
            </PaginationLink>
          </PaginationItem>
        ))}

        <PaginationItem>
          <PaginationNext
            onClick={() => currentPage < totalPages && onPageChange(currentPage + 1)}
            className={
              currentPage === totalPages ? "pointer-events-none opacity-50 cursor-not-allowed" : "cursor-pointer"
            }
          />
        </PaginationItem>
      </PaginationContent>
    </Pagination>
  );
};

export default function ClustersPage() {
  const { clusters, totalItems, loading, page, setPage, fetchClusters } =
    useStore();

  useEffect(() => {
    fetchClusters();  // Fetch all once
  }, [fetchClusters]);

  const totalPages = totalItems ? Math.ceil(totalItems / ITEMS_PER_PAGE) : 1;

  // Slice current page data
  const paginatedClusters = useMemo(() => {
    const startIdx = (page - 1) * ITEMS_PER_PAGE;
    const endIdx = startIdx + ITEMS_PER_PAGE;
    return clusters.slice(startIdx, endIdx);
  }, [clusters, page]);

  return (
    <div className="p-10 space-y-6">
      <div className="flex justify-between">
        <h1 className="text-3xl font-bold">{`${totalItems} Clusters found`}</h1>
        <ClusterCompareDialog clusters={clusters} />
      </div>

      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        {loading
          ? Array.from({ length: ITEMS_PER_PAGE }).map((_, i) => (
              <Card key={i}>
                <CardHeader>
                  <Skeleton className="h-6 w-32 mb-2" />
                  <Skeleton className="h-4 w-48" />
                </CardHeader>
                <CardContent>
                  <Skeleton className="h-24 w-full" />
                </CardContent>
              </Card>
            ))
          : paginatedClusters.map((cluster) => (
              <ClusterCard key={cluster.clusterID} cluster={cluster} />
            ))}
      </div>

      {totalPages > 1 && (
        <ClustersPagePagination
          currentPage={page}
          totalPages={totalPages}
          onPageChange={(p) => setPage(p)}
        />
      )}
    </div>
  );
}
