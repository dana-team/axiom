import { Button } from "@/components/ui/button";
import { Link } from "react-router-dom";
import NotFoundImg from "@/assets/gopher.svg";

export default function NotFound() {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-background text-foreground text-center">
      <img src={NotFoundImg} alt="404 Not Found" className="w-64 mb-8" />
      <h1 className="text-4xl md:text-5xl font-bold mb-4">
        Well, This Is Awkward...
      </h1>
      <p className="text-lg text-muted-foreground mb-6">
        We couldn't find the page you're looking for.
      </p>
      <Link to="/">
        <Button variant="default" className="text-white">
          Back to Home
        </Button>
      </Link>
    </div>
  );
}
